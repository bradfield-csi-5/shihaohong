#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>

// arbitrarily set constants
#define MAXLINE 2048
#define MAXARGS 128

void eval(char *cmdline);
void parseline(char *buf, char **argv);
void unix_error(char *msg);

// "Skeleton" taken from CSAPP: https://github.com/mofaph/csapp/blob/master/code/ecf/shellex.c
// Goal: Use skeleton and add pieces incrementally to understand overall structure of the program.
// Eventually, see if there are pieces I'd want to customize/implement
int main () {
    char cmdline[MAXLINE];

    while (1) {
        printf("> ");
        fgets(cmdline, MAXLINE, stdin);
        if (feof(stdin)) {
            printf("Goodbye~\n");
            exit(0);
        }

        eval(cmdline);
    }
}

void eval(char *cmdline) {
    char *argv[MAXARGS]; /* Argument list execve() */
    char buf[MAXLINE];   /* Holds modified command line */
    pid_t pid;           /* Process id */

    strcpy(buf, cmdline);
    parseline(buf, argv);

    if (argv[0] == NULL) {
	    return;   /* Ignore empty lines */
    }

	if ((pid = fork()) == 0) {   /* Child runs user job */
	    if (execve(argv[0], argv, NULL) < 0) {
            printf("%s: Command not found.\n", argv[0]);
            exit(0);
	    }

	    int status;
	    if (waitpid(pid, &status, 0) < 0) {
		    unix_error("waitfg: waitpid error");
        }
	}
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

void unix_error(char *msg) {
    fprintf(stderr, "%s: %s\n", msg, strerror(errno));
    exit(0);
}
