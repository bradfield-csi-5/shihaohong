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

data_t dotproduct_reduce_proc_call(vec_ptr u, vec_ptr v) {
   data_t sum = 0, u_val, v_val;
   long len = vec_length(u);
   for (long i = 0; i < len; i++) { // we can assume both vectors are same length
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum += u_val * v_val;
   }

   return sum;
}

data_t dotproduct_unrolled_2_1(vec_ptr u, vec_ptr v) {
   data_t sum = 0, u_val, v_val;
   long len = vec_length(u);
   long limit = len-1;
   long i = 0;
   for (; i < limit; i += 2) { // we can assume both vectors are same length
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum += u_val * v_val;

      get_vec_element(u, i + 1, &u_val);
      get_vec_element(v, i + 1, &v_val);
      sum += u_val * v_val;
   }

   for (; i < len ; i++ ) {
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum += u_val * v_val;
   }

   return sum;
}

data_t dotproduct_unrolled_2_2(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, u_val1, v_val1, u_val2, v_val2;
   long len = vec_length(u);
   long limit = len-1;
   long i = 0;
   for (; i < limit; i += 2) { // we can assume both vectors are same length
      get_vec_element(u, i, &u_val1);
      get_vec_element(v, i, &v_val1);
      sum1 += u_val1 * v_val1;

      get_vec_element(u, i + 1, &u_val2);
      get_vec_element(v, i + 1, &v_val2);
      sum2 += u_val2 * v_val2;
   }

   for (; i < len ; i++ ) {
      get_vec_element(u, i, &u_val1);
      get_vec_element(v, i, &v_val1);
      sum1 += u_val1 * v_val1;
   }

   return sum1 + sum2;
}
