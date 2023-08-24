#include "vec.h"

data_t dotproduct(vec_ptr u, vec_ptr v) {
   data_t sum = 0, u_val, v_val;
   for (long i = 0; i < vec_length(u); i++) { // we can assume both vectors are same length
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum += u_val * v_val;
   }

   return sum;
}

data_t dotproduct_reduce_len_call(vec_ptr u, vec_ptr v) {
   data_t sum = 0, u_val, v_val;
   long len = vec_length(u);
   for (long i = 0; i < len; i++) { // we can assume both vectors are same length
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum += u_val * v_val;
   }

   return sum;
}

data_t dotproduct_reduce_proc_call(vec_ptr u, vec_ptr v) {
   data_t sum = 0;
   data_t *u_start, *v_start;
   long len = vec_length(u);
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (long i = 0; i < len; i++) { // we can assume both vectors are same length
      sum += u_start[i] * v_start[i];
   }

   return sum;
}


data_t dotproduct_unrolled_2_1(vec_ptr u, vec_ptr v) {
   data_t sum = 0;
   long len = vec_length(u);
   long limit = len-1;
   long i = 0;
   data_t *u_start, *v_start;
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (; i < limit; i += 2) { // we can assume both vectors are same length
      sum += u_start[i] * v_start[i];
      sum += u_start[i + 1] * v_start[i + 1];
   }

   for (; i < len ; i++ ) {
      sum += u_start[i] * v_start[i];
   }

   return sum;
}

data_t dotproduct_unrolled_2_2(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0;
   long len = vec_length(u);
   long limit = len-1;
   long i = 0;
   data_t *u_start, *v_start;
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (; i < limit; i += 2) { // we can assume both vectors are same length
      sum1 += u_start[i] * v_start[i];
      sum2 += u_start[i + 1] * v_start[i + 1];
   }

   for (; i < len ; i++ ) {
      sum1 += u_start[i] * v_start[i];
   }

   return sum1 + sum2;
}

data_t dotproduct_unrolled_4_4(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, sum3 = 0, sum4 = 0;
   long len = vec_length(u);
   long limit = len-3;
   long i = 0;
   data_t *u_start, *v_start;
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (; i < limit; i += 4) { // we can assume both vectors are same length
      sum1 += u_start[i] * v_start[i];
      sum2 += u_start[i + 1] * v_start[i + 1];
      sum3 += u_start[i + 2] * v_start[i + 2];
      sum4 += u_start[i + 3] * v_start[i + 3];
   }

   for (; i < len ; i++ ) {
      sum1 += u_start[i] * v_start[i];
   }

   return sum1 + sum2 + sum3 + sum4;
}

data_t dotproduct_unrolled_6_6(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, sum3 = 0, sum4 = 0, sum5 = 0, sum6 = 0;
   long len = vec_length(u);
   long limit = len-5;
   long i = 0;
   data_t *u_start, *v_start;
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (; i < limit; i += 6) { // we can assume both vectors are same length
      sum1 += u_start[i] * v_start[i];
      sum2 += u_start[i + 1] * v_start[i + 1];
      sum3 += u_start[i + 2] * v_start[i + 2];
      sum4 += u_start[i + 3] * v_start[i + 3];
      sum5 += u_start[i + 4] * v_start[i + 4];
      sum6 += u_start[i + 5] * v_start[i + 5];
   }

   for (; i < len ; i++ ) {
      sum1 += u_start[i] * v_start[i];
   }

   return sum1 + sum2 + sum3 + sum4 + sum5 + sum6;
}

data_t dotproduct_unrolled_8_8(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, sum3 = 0, sum4 = 0, sum5 = 0, sum6 = 0;
   data_t sum7 = 0, sum8 = 0;
   long len = vec_length(u);
   long limit = len-7;
   long i = 0;
   data_t *u_start, *v_start;
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (; i < limit; i += 8) { // we can assume both vectors are same length
      sum1 += u_start[i] * v_start[i];
      sum2 += u_start[i + 1] * v_start[i + 1];
      sum3 += u_start[i + 2] * v_start[i + 2];
      sum4 += u_start[i + 3] * v_start[i + 3];
      sum5 += u_start[i + 4] * v_start[i + 4];
      sum6 += u_start[i + 5] * v_start[i + 5];
      sum7 += u_start[i + 6] * v_start[i + 6];
      sum8 += u_start[i + 7] * v_start[i + 7];
   }

   for (; i < len; i++) {
      sum1 += u_start[i] * v_start[i];
   }

   return sum1 + sum2 + sum3 + sum4 + sum5 + sum6 + sum7 + sum8;
}


data_t dotproduct_unrolled_10_10(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, sum3 = 0, sum4 = 0, sum5 = 0, sum6 = 0;
   data_t sum7 = 0, sum8 = 0, sum9 = 0, sum10 = 0;
   long len = vec_length(u);
   long limit = len-9;
   long i = 0;
   data_t *u_start, *v_start;
   u_start = get_vec_start(u);
   v_start = get_vec_start(v);
   for (; i < limit; i += 10) { // we can assume both vectors are same length
      sum1 += u_start[i] * v_start[i];
      sum2 += u_start[i + 1] * v_start[i + 1];
      sum3 += u_start[i + 2] * v_start[i + 2];
      sum4 += u_start[i + 3] * v_start[i + 3];
      sum5 += u_start[i + 4] * v_start[i + 4];
      sum6 += u_start[i + 5] * v_start[i + 5];
      sum7 += u_start[i + 6] * v_start[i + 6];
      sum8 += u_start[i + 7] * v_start[i + 7];
      sum9 += u_start[i + 8] * v_start[i + 8];
      sum10 += u_start[i + 9] * v_start[i + 9];
   }

   for (; i < len ; i++ ) {
      sum1 += u_start[i] * v_start[i];
   }

   return sum1 + sum2 + sum3 + sum4 + sum5 + sum6 + sum7 + sum8 + sum9 + sum10;
}
