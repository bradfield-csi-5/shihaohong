#include <errno.h>
#include <signal.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <sys/wait.h>
#include <unistd.h>

#define BASE_SOCK_PATH "socket"

int START = 2, END = 20;
char *TESTS[] = {"brute_force", "brutish", "miller_rabin"};
int num_tests = sizeof(TESTS) / sizeof(char *);

int main(int argc, char *argv[]) {
  int s, len;
  struct sockaddr_un remote = {
    .sun_family = AF_UNIX,
    // .sun_path = SOCK_PATH,   // Can't do assignment to an array
  };

  int result, i;
  long n;
  pid_t pid;

  // printf("num tests: %d\n", num_tests);

  for (i = 0; i < 1; i++) {
  // for (i = 0; i < num_tests; i++) {

    // TODO: get fork to work?
    // pid = fork();

    // if (pid == -1) {
    //   fprintf(stderr, "Failed to fork\n");
    //   exit(-1);
    // }

    // if (pid == 0) {
    //   execl("sprimality", "sprimality", TESTS[i], (char *)NULL);
    // }

    // we are the parent
    if ((s = socket(AF_UNIX, SOCK_STREAM, 0)) == -1) {
      perror("socket");
      exit(1);
    }

    printf("Trying to connect...\n");

    char path[100];
    strcpy(path, BASE_SOCK_PATH);
    strcat(path, "_");
    strcat(path, TESTS[i]);
    strcpy(remote.sun_path, path);
    printf("Socket: %s\n", path);
    len = strlen(remote.sun_path) + sizeof(remote.sun_family) + 1;
    if (connect(s, (struct sockaddr *)&remote, len) == -1) {
      perror("connect");
      exit(1);
    }
    printf("%s has connected.\n", path);
  }

  // for each number, run each test
  for (n = START; n <= END; n++) {
    for (i = 0; i < 1; i++) {
    // for (i = 0; i < num_tests; i++) {
      char input[10];
      snprintf(input, sizeof(input), "%ld", n);
      printf("sending input: %s\n", input);

      if (send(s, input, strlen(input)+1, 0) == -1) {
        perror("send");
        exit(1);
      }

      char result[100];
      result[0] = '\0';
      printf("receiving results:\n");
      if ((len=recv(s, result, sizeof(result), 0)) > 0) {
        result[len] = '\0';
        printf("num:res => %s\n", result);
      } else {
        if (len < 0) perror("recv");
        else printf("Server closed connection\n");
        exit(1);
      }
    }
  }

  // long received = 0;
  // for (;received < END;) {
  //   char result[100];
  //   result[0] = '\0';
  //   printf("receiving results:\n");
  //   if ((len=recv(s, result, sizeof(result), 0)) > 0) {
  //     result[len] = '\0';
  //     printf("echo> %s\n", result);
  //     received++;
  //   } else {
  //     if (len < 0) perror("recv");
  //     else printf("Server closed connection\n");
  //     exit(1);
  //   }
  // }

  close(s);
  return 0;
}
