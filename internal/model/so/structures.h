#ifndef SRC_CALCULATOR_STRUCTURES_H_
#define SRC_CALCULATOR_STRUCTURES_H_

#include <math.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

/// FUNCTIONS_OUTPUT
#define SUCCESS 0
#define FAILURE 1

/// STACK
typedef struct node {
  double value;
  char symbol;
  struct node *next;
} node_t;

/// FOR PARSING
struct tokens {
  int8_t i;
  int8_t values;
  int8_t symbols;
  int8_t brackets;
  char temp_symbols;
  double temp_res;
  struct operators *operator_before_brackets;
};

typedef struct operators {
  char symbol;
  int8_t brackets;
  struct operators *next;
} operators_t;

#endif  // SRC_CALCULATOR_STRUCTURES_H_
