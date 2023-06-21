#include <stdio.h>
#include <sys/types.h>
#include <dirent.h>

#define MAX_PATH 256

// See https://www.gnu.org/software/libc/manual/html_node/Accessing-Directories.html
// for the directory access APIs.

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
    // TODO: filter results by hidden files (ignore files that start with ".")
    dp = opendir(path);
    if (dp != NULL) {
        // reads directory stream until NULL when error/end
        while ((ep = readdir(dp))) {
            // TODO: sort like the actual `ls` program would (sort dir and non-dir lexicographically, separately)
            // TODO: display more directory information
            printf("%s\n", ep->d_name);
        }
        closedir(dp);
    } else {
        printf("Couldn't open the directory\n");
    }

    return 0;
}
