#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <math.h>
#include <omp.h>

#define N 100000
#define num_threads 5


// Merge函数合并两个子数组形成单一的已排好序的字数组
// 并代替当前的子数组A[p..r]
void Merge(int *a, int p, int q, int r) {
    int res[r-p+1];
    int ci = 0;

    int li = p;
    int ri = q+1;

    while (li <= q && ri <= r) {
        if (a[li] < a[ri]) {
            res[ci] = a[li++];
        } else {
            res[ci] = a[ri++];
        }
        ci++;
    }
    while (li <= q) res[ci++] = a[li++];
    while (ri <= r) res[ci++] = a[ri++];

    int ti = r;
    while (--ci >= 0) a[ti--] = res[ci]; 
}

//归并排序
void MergeSort(int *a, int p, int r) {
    if (p >= r) return;

    int q = (p + r) / 2;
    MergeSort(a, p, q);
    MergeSort(a, q+1, r);
    Merge(a, p, q, r);
}

int min(int a, int b) {
    if (a < b) return a;
    return b;
}

int id;
int base = int(ceil(n * 1. / num_threads));  //划分的每段段数
int sample_len = 0;
int samples[num_threads * num_threads];  //划分的每段段数
int pivot[num_threads - 1];  //主元
int count[num_threads];  //每个处理器每段的个数
int pivot_array[num_threads][N];  //处理器数组空间
void PSRS(int *array, int n) {
    omp_set_num_threads(num_threads);
    #pragma omp parallel shared(base, array, n, sample_len, samples, pivot, count) private(id)
    {
        id = omp_get_thread_num();

        int lower = id * base;
        int upper = min(n - 1, (id + 1) * base - 1);

        //每个处理器对所在的段进行局部串行归并排序
        MergeSort(array, lower, upper-1);

        #pragma omp critical
        //每个处理器选出P个样本，进行正则采样
        for (int k = 0; k < num_threads; k++)
            samples[sample_len++] = array[lower + base / num_threads * k];
        
        //设置路障，同步队列中的所有线程
        #pragma omp barrier
        
        //主线程对采样的p个样本进行排序
        #pragma omp master
        {
            MergeSort(samples, 0, sample_len - 1);
            //选出num_threads-1个主元
            for (int m = 0; m < num_threads - 1; m++)
                pivot[m] = samples[(m + 1) * num_threads];
        }
        
        #pragma omp barrier
        
        //根据主元对每一个cpu数据段进行划分
        count[id] = 0;
        for (int k = 0; k < num_threads; k++) {
            int lower = k * base;
            int upper = min((k+1)*base-1, n-1);
            for (int ai = lower; ai < upper; ai++) {
                if (id == 0) {
                    if (array[ai] >= pivot[0]) break; 
                    pivot_array[id][count[id]++] = array[ai];
                
                } else if (id == num_threads - 1) {
                    if (array[ai] < pivot[id-1]) continue;
                    pivot_array[id][count[id]++] = array[ai];
                
                } else {
                    if (array[ai] < pivot[id-1]) continue;
                    if (array[ai] >= pivot[id]) break;
                    pivot_array[id][count[id]++] = array[ai];
                }
            }
        }

        MergeSort(pivot_array[id], 0, count[id]-1);

        #pragma omp barrier

        int pre_count = 0;
        for (int k = 0; k < num_threads; k++) {
            if (k < id) {
                pre_count += count[k];
                continue;
            }

            memcpy(array + pre_count, pivot_array[id], sizeof(int)*count[id]);
            break;
        }
    }
    
    //向各个线程发送数据，各个线程自己排序
    // #pragma omp parallel shared(pivot_array, count)
    // {
    //     int id = omp_get_thread_num();
    //     for (int k = 0; k < num_threads; k++)
    //     {
    //         if (k != id)
    //         {
    //             memcpy(pivot_array[id][id] + count[id][id], pivot_array[k][id], sizeof(int) * count[k][id]);
    //             count[id][id] += count[k][id];
    //         }
    //     }
    //     MergeSort(pivot_array[id][id], 0, count[id][id] - 1);
    // }

    // #pragma omp barrier
    // printf("result:\n");
    // for (int k = 0; k < n; k++) printf("%d ", array[k]);
    // for (int k = 0; k < num_threads; k++)
    // {
    //     for (int m = 0; m < count[k]; m++)
    //     {
    //         printf("%d ", pivot_array[k][m]);
    //     }
    //     printf("\n");
    // }
}

int array[N];

int main() {
    for (int i = 0; i < N; i++) array[i] = rand();

    PSRS(array, N);
    // MergeSort(array, 0, 35);
    // for(int i = 0; i < 36; i++)
    // {
    //     printf("%d ", array[i]);
    // }
    return 0;
}