#include "vec.h"


data_t dotproduct(vec_ptr u, vec_ptr v) {
   data_t sum1 = 0, sum2 = 0, u_val, v_val;
   long len = vec_length(u);
   long limit = len-1;
   long i = 0;
   for (; i < limit; i += 2) { // we can assume both vectors are same length
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum1 += u_val * v_val;

      get_vec_element(u, i + 1, &u_val);
      get_vec_element(v, i + 1, &v_val);
      sum2 += u_val * v_val;
   }

   for (; i < len ; i++ ) {
      get_vec_element(u, i, &u_val);
      get_vec_element(v, i, &v_val);
      sum1 += u_val * v_val;
   }

   return sum1 + sum2;
}
