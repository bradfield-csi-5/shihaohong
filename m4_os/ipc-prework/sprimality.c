#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <unistd.h>

#define BASE_SOCK_PATH "socket"

int brute_force(long n);
int brutish(long n);
int miller_rabin(long n);

void exit_with_usage() {
  fprintf(stderr, "Usage: ./sprimality [brute_force|brutish|miller_rabin]\n");
  exit(1);
}

int main(int argc, char*argv[]) {
  long num;
  int (*func)(long);

  if (argc != 2)
    exit_with_usage();

  char path[200];
  strcpy(path, BASE_SOCK_PATH);
  if (strcmp(argv[1], "brute_force") == 0) {
    func = &brute_force;
    strcat(path, "_brute_force");
  } else if (strcmp(argv[1], "brutish") == 0) {
    func = &brutish;
    strcat(path, "_brutish");
  } else if (strcmp(argv[1], "miller_rabin") == 0) {
    func = &miller_rabin;
    strcat(path, "_miller_rabin");
  } else {
    exit_with_usage();
  }

  int s, s2, len;
  struct sockaddr_un remote, local = {
    .sun_family = AF_UNIX,
    // .sun_path = ,  // Can't do assignment to an array
  };

  if ((s = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
    perror("socket");
    exit(1);
  }

  strcpy(local.sun_path, path);
  printf("path: %s\n", local.sun_path);
  unlink(local.sun_path);

  // TODO: figure out why there's off-by-one error in the len
  len = strlen(local.sun_path) + sizeof(local.sun_family) + 1;
  if (bind(s, (struct sockaddr *)&local, len) == -1) {
    perror("bind");
    exit(1);
  }

  // The second argument is the number of incoming connections that can be queued before you call accept()
  if (listen(s, 5) == -1) {
    perror("listen");
    exit(1);
  }

  char str[100];
  for(;;) {
    int done, n;
    printf("Waiting for a connection...\n");
    socklen_t slen = sizeof(remote);
    if ((s2 = accept(s, (struct sockaddr *)&remote, &slen)) == -1) {
      perror("accept");
      exit(1);
    }

    printf("Connected.\n");

    done = 0;
    do {
      n = recv(s2, str, sizeof(str), 0);
      if (n <= 0) {
        if (n < 0) perror("recv");
        done = 1;
      }

      if (!done) {
        long result = (*func)(atol(str));
        char res[5];
        sprintf(res,"%ld",result);
        char output[100];
        output[0] = '\0';
        strcat(output, str);
        strcat(output, ":");
        strcat(output, res);

        printf("returned val: %s\n", output);

        if (send(s2, output, sizeof(output), 0) < 0) {
          perror("send");
          done = 1;
        }
      }
    } while (!done);

    close(s2);
  }

  return 0;
}

/*
 * Primality test implementations
 */

// Just test every factor
int brute_force(long n) {
  for (long i = 2; i < n; i++)
    if (n % i == 0)
      return 0;
  return 1;
}

// Test factors, up to sqrt(n)
int brutish(long n) {
  long max = floor(sqrt(n));
  for (long i = 2; i <= max; i++)
    if (n % i == 0)
      return 0;
  return 1;
}

int randint(int a, int b) { return rand() % (++b - a) + a; }

int modpow(int a, int d, int m) {
  int c = a;
  for (int i = 1; i < d; i++)
    c = (c * a) % m;
  return c % m;
}

int witness(int a, int s, int d, int n) {
  int x = modpow(a, d, n);
  if (x == 1)
    return 1;
  for (int i = 0; i < s - 1; i++) {
    if (x == n - 1)
      return 1;
    x = modpow(x, 2, n);
  }
  return (x == n - 1);
}

// TODO we should probably make this a parameter!
int MILLER_RABIN_ITERATIONS = 10;

// An implementation of the probabilistic Miller-Rabin test
int miller_rabin(long n) {
  int a, s = 0, d = n - 1;

  if (n == 2)
    return 1;

  if (!(n & 1) || n <= 1)
    return 0;

  while (!(d & 1)) {
    d >>= 1;
    s += 1;
  }
  for (int i = 0; i < MILLER_RABIN_ITERATIONS; i++) {
    a = randint(2, n - 1);
    if (!witness(a, s, d, n))
      return 0;
  }
  return 1;
}
