#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <setjmp.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>

// arbitrarily set constants
#define MAXLINE 2048
#define MAXARGS 128

volatile sig_atomic_t pid;
jmp_buf j_buf;

void eval(char *cmdline, sigset_t *mask, sigset_t *prev);
void parseline(char *buf, char **argv);
void unix_error(char *msg);
void sigchld_handler(int s);
void sigint_handler(int s);

typedef void handler_t(int);
handler_t *Signal(int signum, handler_t *handler);

// "Skeleton" taken from CSAPP: https://github.com/mofaph/csapp/blob/master/code/ecf/shellex.c
// Goal: Use skeleton and add pieces incrementally to understand overall structure of the program.
// Eventually, see if there are pieces I'd want to customize/implement
int main () {
    char cmdline[MAXLINE];
    sigset_t mask, prev;
    Signal(SIGCHLD, sigchld_handler);
    Signal(SIGINT, sigint_handler);
    sigemptyset(&mask);
    sigaddset(&mask, SIGCHLD);

    while (1) {
        setjmp(j_buf);
        printf("> ");
        fgets(cmdline, MAXLINE, stdin);
        if (feof(stdin)) {
            printf("Goodbye~\n");
            exit(0);
        }

        eval(cmdline, &mask, &prev);
    }
}

void eval(char *cmdline, sigset_t *mask, sigset_t *prev) {
    char *argv[MAXARGS]; /* Argument list execve() */
    char buf[MAXLINE];   /* Holds modified command line */

    strcpy(buf, cmdline);
    parseline(buf, argv);

    if (argv[0] == NULL) {
	    return;   /* Ignore empty lines */
    }

    sigprocmask(SIG_BLOCK, mask, prev);
	if ((pid = fork()) == 0) {   /* Child runs user job */
	    if (execve(argv[0], argv, NULL) < 0) {
            printf("%s: Command not found.\n", argv[0]);
            exit(0);
	    }
	}

    while (!pid) {
        sigsuspend(prev);
    }
    sigprocmask(SIG_SETMASK, prev, NULL); /* Unblock SIGCHLD */
}

void parseline(char *buf, char **argv) {
    char *delim;         /* Points to first space delimiter */
    int argc;            /* Number of args */

    buf[strlen(buf)-1] = ' ';  /* Replace trailing '\n' with space */
    while (*buf && (*buf == ' ')) {
	    buf++;
    }

    /* Build the argv list */
    argc = 0;
    while ((delim = strchr(buf, ' '))) {
        // set delim to \0 and then save
        // the argument into the argv array
        *delim = '\0';
        argv[argc] = buf;
        argc += 1;

        // iterate forwards
        buf = delim + 1;

        // Ignore spaces
        while (*buf && (*buf == ' ')) {
            buf++;
        }
    }
    argv[argc] = NULL;

    if (argc == 0) {
    	return;
    }
}

handler_t *Signal(int signum, handler_t *handler) {
    struct sigaction action, old_action;

    action.sa_handler = handler;
    sigemptyset(&action.sa_mask); /* block sigs of type being handled */
    action.sa_flags = SA_RESTART; /* restart syscalls if possible */

    if (sigaction(signum, &action, &old_action) < 0)
        unix_error("Signal error");
    return (old_action.sa_handler);
}

void sigchld_handler(int s) {
    int olderrno = errno;
    pid = waitpid(-1, NULL, 0);
    errno = olderrno;
}

void sigint_handler(int s) {
    if (pid == 0) {
        exit(0);
    }
    printf("\n");
    longjmp(j_buf, 0);
}

void unix_error(char *msg) {
    fprintf(stderr, "%s: %s\n", msg, strerror(errno));
    exit(0);
}
