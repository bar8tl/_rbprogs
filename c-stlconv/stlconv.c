/*
stlconv.c [2010-03-01 BAR8TL]
Convert column formatted file into a tabs formatted file
*/
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define BUFSIZE 256
#define PRMSIZE 65
#define NCOLS   23

int   convFile (char*, char*, char*);
char *substr   (char*, char*, unsigned int, int);
char *trim     (char*, char*);
int   blanks   (char);
char *ufmtMatnr(char*, char*);

struct layout {
  char *fld;
  int pos, len;
} column[] = {
  "werks",   1,  4,
  "matnr",   6, 18,
  "maktx",  25, 40,
  "meins",  66,  3,
  "spart",  70,  2,
  "matkl",  73,  9,
  "prctr",  83, 10,
  "ekgrp",  94,  3,
  "dispo",  98,  3,
  "maabc", 102,  1,
  "mtart", 104,  4,
  "beskz", 109,  1,
  "sobsl", 111,  2,
  "kzumw", 114,  1,
  "mhdhb", 116,  5,
  "plifz", 122,  3,
  "rwpro", 126,  3,
  "mrppp", 130,  3,
  "bstrf", 134, 16,
  "mein2", 152,  3,
  "xchpf", 156,  1,
  "ersda", 158, 10,
  "prdha", 169, 18
};

int main(int argc, char **argv) {
  char src[PRMSIZE] = "mtrlmstr.txt";
  char tgt[PRMSIZE] = "mtlmstr.txt";
  char plt[PRMSIZE] = "6970";
  if (argc >= 1)
    strcpy(src, argv[1]);
  if (argc >= 2)
    strcpy(tgt, argv[2]);
  if (argc >= 3)
    strcpy(plt, argv[3]);
  convFile(src, tgt, plt);
  return(0);
}

int convFile(char *src, char *tgt, char *plt) {
  FILE *ifile;
  FILE *ofile;
  int i;
  char s[BUFSIZE];
  char r[BUFSIZE];
  char t[BUFSIZE];
  char l[BUFSIZE];

  ifile = fopen(src, "r");
  if (!ifile) {
    printf("File %s cannot be opened\n", src);
    exit(1);
  }
  ofile = fopen(tgt, "w");
  fgets(s, BUFSIZE, ifile);
  while (!feof(ifile)) {
    if (strcmp(substr(s, r, column[0].pos, column[0].len), plt) == 0) {
      r[0] = l[0] = t[0] = '\0';
      strcat(l, substr(s, r, column[0].pos, column[0].len));
      for (i = 1; i < NCOLS; i++) {
        strcat(l, "\t");
        if (i == 1)
          strcat(l, ufmtMatnr(substr(s, r, column[i].pos, column[i].len), t));
        else
          strcat(l, trim(substr(s, r, column[i].pos, column[i].len), t));
      }
      strcat(l, "\t");
      strcat(l, trim(substr(s, r, column[1].pos, column[1].len), t));
      strcat(l, "\n");
      fputs(l, ofile);
    }
    fgets(s, BUFSIZE, ifile);
  }
  fclose(ifile);
  fclose(ofile);
  return(0);
}

char *substr(char* s1, char* s2, unsigned int offst, int lngth) {
  int i;

  *s2 = '\0';
  if (offst < strlen(s1)) {
    for (i = 0; i < lngth && *(s1+offst+i) != '\0'; i++)
      *(s2+i) = *(s1+offst+i);
    *(s2+i) = '\0';
  }
  return s2;
}

char *trim(char *s1, char *s2) {
  int i, j;

  for (i = j = 0; s1[i] != '\0'; i++)
    if (!blanks(s1[i]) || j > 0)
      s2[j++] = *(s1+i);
  s2[j] = '\0';
  for (i = strlen(s1)-1; blanks(s2[i]); i--)
    s2[i] = '\0';
  return s2;
}

int blanks(char c) {
  return (c == ' ' || c == '\t' || c == '\r' || c == '\n' || c == '\0' ||
    c == '\a' || c == '\b' || c == '\f' || c == '\v');
}

char *ufmtMatnr(char *s1, char *s2) {
  int i, j;

  for (i = j = 0; s1[i] != '\0'; i++)
    if (s1[i] != '.' && s1[i] != '-' && s1[i] != '/' && s1[i] != ' ')
      s2[j++] = *(s1+i);
  s2[j] = '\0';
  return s2;
}
