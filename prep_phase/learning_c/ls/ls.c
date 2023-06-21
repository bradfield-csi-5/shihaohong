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

int main(int argc, char *argv[]) {
    DIR *dp;
    struct dirent *ep;

    char* path;
    char* last_arg = argv[argc - 1];

    int a_option = 0;

    // look for and process options first
    // once non-option found, assume it's filepaths
    // does not process "--" for now
    int i;
    for (i = 1; i < argc; i++) {
        char* arg = argv[i];
        char first_char = arg[0];
        if (first_char == '-') {
            // process options
            if (strstr(arg, "a")) {
                a_option = 1;
            }
        } else {
            break;
        }
    }

    // for now, assume only one file processed
    // TODO: handle multiple file directory inputs in argument
    // ie. ls ./ ../ (should list the contents of both directories,
    // separated by \n\n). For now, assume last arg is the dir to be listed
    if (argc < 2 || i == argc) {
        path = "./";
    } else {
        path = last_arg;
    }

    // if somehow NULL, default to pwd
    if (path == NULL) {
        path = "./";
    }

    // opens a directory stream
    dp = opendir(path);
    if (dp != NULL) {
        // reads directory stream until NULL when error/end
        while ((ep = readdir(dp))) {
            // skip if hidden file and "a" flag present
            if (!a_option && ep->d_name[0] == '.')
                continue;

            if (ep->d_type == DT_DIR || ep->d_type == DT_REG) {
                dir[i_d] = *ep;
                i_d++;
            } // TODO: do i need to handle other file types?
        }
        closedir(dp);
    } else {
        printf("Couldn't open the directory\n");
        return CNOD_ERR;
    }

    qsort(dir, i_d, sizeof(struct dirent), dircmp);
    dirlist();
    return 0;
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
