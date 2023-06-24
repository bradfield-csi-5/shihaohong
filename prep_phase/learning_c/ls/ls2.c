#include <stdio.h>
#include <sys/types.h>
#include <dirent.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

// arbitrarily set for now
#define MAX_FILES 256
#define FILENAME_LENGTH 256

struct dircontents {
    char filename[FILENAME_LENGTH];
    struct dirent dirents[MAX_FILES];
    int numdirs;
    int couldread;
};

int main(int argc, char* argv[]) {
    // parse options
    int a_option = 0;
    char c;
    while ((c = getopt(argc, argv, "a:")) != -1) {
        switch (c) {
            case 'a':
                a_option = 1;
                break;
        }
    }
    printf("a: %d\n", a_option);

    int n_files;
    char* arg;
    DIR *dp;
    struct dirent *ep;

    // first pass so can get num files
    if (argc == 1) {
        n_files = 0;
    } else {
        for (n_files = 0; n_files < argc - 1; n_files++) {
            arg = argv[argc - n_files - 1];
            // stop when reach options
            if (arg[0] == '-') {
                break;
            }
        }
    }

    struct dircontents *dirs;
    // if no args found, set to 1 and default path to "/."
    if (n_files == 0) {
        n_files = 1;
        dirs = (struct dircontents *) malloc(sizeof(struct dircontents));
        dp = opendir("./");
        if (dp != NULL) {
            strcpy(dirs[0].filename, "./");
            dirs[0].numdirs = 0;
            dirs[0].couldread = 1;

            while ((ep = readdir(dp))) {
                if (ep->d_type == DT_DIR || ep->d_type == DT_REG) {
                    if (a_option != 1 && ep->d_name[0] == '.') {
                        continue;
                    }

                    // save to dir list
                    dirs[0].dirents[dirs[0].numdirs] = *ep;
                    dirs[0].numdirs++;

                } // TODO: do i need to handle other file types?
            }
            (void) closedir(dp);
        } else {
            // TODO: failure case
        }
    } else {
        dirs = (struct dircontents *) malloc(sizeof(struct dircontents) * n_files);

        for (int i = 0; i < n_files; i++) {
            arg = argv[argc - i - 1];
            dp = opendir(arg);
            if (dp != NULL) {
                strcpy(dirs[i].filename, arg);
                dirs[i].numdirs = 0;
                dirs[i].couldread = 1;

                while ((ep = readdir(dp))) {
                    if (ep->d_type == DT_DIR || ep->d_type == DT_REG) {
                        if (a_option != 1 && ep->d_name[0] == '.') {
                            continue;
                        }

                        // save to dir list
                        dirs[i].dirents[dirs[i].numdirs] = *ep;
                        dirs[i].numdirs++;

                    } // TODO: do i need to handle other file types?
                }
                (void) closedir(dp);
            } else {
                // TODO: failure case
            }
        }
    }

    // TODO: cleanup print logic
    for (int i = 0; i < n_files; i++) {
        struct dircontents *current = &dirs[i];
        printf("current file: %s\n", current->filename);
        int num_filents = current->numdirs;
        for (int j = 0; j < num_filents; j++) {
            // TODO: sort print output lexicographically
            printf("%s\n", current->dirents[j].d_name);
        }
        printf("\n");
    }
}
