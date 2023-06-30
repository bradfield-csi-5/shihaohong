#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <dirent.h>

#define MAX_DIR 65536

// See https://www.gnu.org/software/libc/manual/html_node/Accessing-Directories.html
// for the directory access APIs.

int dircmp(const void *a, const void *b);
void dirlist(struct dirent dir[MAX_DIR], int len, char* paths);
struct dirent dir[MAX_DIR];
int i_d = 0;

// TODO: `ls` lists directory does not exist at the beginning before listing any
// real files/directories.
// TODO: `ls` does not prefix with "$path:\n" if there's only one path to process
// TODO: probably best to create a struct that houses all the dirs to parse
// and any relevant metadata
int main(int argc, char *argv[]) {
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

    int path_len = ((argc < 2) || (argc = i)) ? 1 : argc - i;
    char* paths[path_len];
    if (argc < 2 || argc == i) {
        paths[0] = "./";
    } else {
        int path_i;
        for (path_i = 0; path_i < path_len; i++, path_i++) {
            paths[path_i] = argv[i];
        }
    }

    // opens a directory stream
    DIR *dp;
    struct dirent *ep;
    for (i = 0; i < path_len; i++) {
        dp = opendir(paths[i]);
        if (dp != NULL) {
            i_d = 0;
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
            qsort(dir, i_d, sizeof(struct dirent), dircmp);
            dirlist(dir, i_d, paths[i]);
        } else {
            printf("%s: No such directory\n\n", paths[i]);
        }
    }

    return 0;
}

int dircmp(const void *a, const void* b) {
    struct dirent* a_d;
    struct dirent* b_d;
    a_d = (struct dirent*) a;
    b_d = (struct dirent*) b;
    return strcmp(a_d->d_name, b_d->d_name);
}

void dirlist(struct dirent dir[MAX_DIR], int len, char* path) {
    printf("%s:\n", path);
    int i;
    struct dirent *d;
    for (i = 0; i < len; i++) {
        d = &dir[i];
        printf("%s\n", d->d_name);
    }
    printf("\n");
}
