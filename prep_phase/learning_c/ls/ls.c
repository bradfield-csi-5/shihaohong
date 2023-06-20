#include <stdio.h>
#include <sys/types.h>
#include <dirent.h>

// See https://www.gnu.org/software/libc/manual/html_node/Accessing-Directories.html
// for the directory access APIs.

// TODO: error handling. exit 0 on success, >0 with errors
int main() {
    DIR *dp;
    struct dirent *ep;

    // opens a directory stream
    // TODO: take argument for dir to read instead of hard-coded current dir
    // handle both relative path and absolute path
    // TODO: filter results by hidden files (ignore files that start with ".")
    dp = opendir("./");
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
