#ifndef SRC_CALCULATOR_CALC_H_
#define SRC_CALCULATOR_CALC_H_

#include "structures.h"

int8_t calc(char *str, double *result);

/// ========================== STACK ===========================
void remove_stack(node_t **stack);
int8_t stack_rewind(node_t **stacks, struct tokens *);
int8_t push(node_t **stack, double value, char);
int8_t pop(node_t **stack, double *value, char *);

/// ======================== PARSE_CALC ========================
int8_t calc_parsing(char *, struct tokens *tokens_t, node_t *);
int8_t smart_push(node_t **stacks, struct tokens *tokens_t);
struct tokens *save_operator(struct tokens *tokens_t);
struct tokens *take_operator(struct tokens *tokens_t);
void get_operator(struct tokens *tokens_t, char *str);
int8_t a_to_double(char *str, struct tokens *tokens_t);
void remove_tokens_t(struct tokens *tokens_t);
int8_t get_double(struct tokens *tokens_t);
int8_t check_unary(struct tokens *);
int8_t get_priority(char);

/// ========================= RUN_CALC =========================
int8_t run_calc(node_t **stacks, struct tokens *tokens_t);
double binary_calc(double *, double *, char);
double unary_calc(double *value, char);

#endif  // SRC_CALCULATOR_CALC_H_
