#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <dirent.h>

#define MAX_PATH 256
#define MAX_DIR 65536

#define CNOD_ERR 1

// See https://www.gnu.org/software/libc/manual/html_node/Accessing-Directories.html
// for the directory access APIs.

struct dirent dir[MAX_DIR];
int i_d = 0;
int dircmp(const void *a, const void *b);
void dirlist();

// TODO: error handling. exit 0 on success, >0 with errors
int main(int argc, char *argv[]) {
    DIR *dp;
    struct dirent *ep;

    // hard-code for first argument, ignore options for now
    // TODO: implement taking options
    char* path = argv[1];
    if (path == NULL) {
        path = "./";
    }

    // opens a directory stream
    dp = opendir(path);
    if (dp != NULL) {
        // reads directory stream until NULL when error/end
        while ((ep = readdir(dp))) {
            if (ep->d_name[0] != '.') {
                if (ep->d_type == DT_DIR || ep->d_type == DT_REG) {
                    dir[i_d] = *ep;
                    i_d++;
                } // TODO: do i need to handle other file types?
            }
        }
        closedir(dp);
    } else {
        printf("Couldn't open the directory\n");
        return CNOD_ERR;
    }

    qsort(dir, i_d, sizeof(struct dirent), dircmp);
    dirlist();
}

int dircmp(const void *a, const void* b) {
    struct dirent* a_d;
    struct dirent* b_d;
    a_d = (struct dirent*) a;
    b_d = (struct dirent*) b;
    return strcmp(a_d->d_name, b_d->d_name);
}

void dirlist() {
    int i;
    struct dirent d;
    for (i = 0; i < i_d; i++) {
        d = dir[i];
        printf("%s\n", d.d_name);
    }
}
