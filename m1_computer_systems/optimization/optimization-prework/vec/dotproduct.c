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

   for (; i < len ; i++ ) {
      sum1 += u_start[i] * v_start[i];
   }

   return sum1 + sum2;
}
