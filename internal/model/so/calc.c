#include "calc.h"

#include "structures.h"

int8_t calc(char *str, double *result) {
  int8_t output = SUCCESS;
  node_t *stack = NULL;
  struct tokens tokens_t = {0};

  output = calc_parsing(!str ? "" : str, &tokens_t, stack);
  output = tokens_t.brackets ? FAILURE : output;
  *result = output ? 0.0 : tokens_t.temp_res;
  remove_tokens_t(&tokens_t);
  remove_stack(&stack);

  return output;
}

int8_t a_to_double(char *str, struct tokens *tokens_t) {
  int8_t output = SUCCESS;

  if (str[tokens_t->i] > 47 && str[tokens_t->i] < 58) {
    while ((str[tokens_t->i] > 47 && str[tokens_t->i] < 58) ||
           str[tokens_t->i] == '.') {
      tokens_t->temp_res = tokens_t->temp_res * 10 + (str[tokens_t->i++] - '0');
      if (str[tokens_t->i] == '.') {
        tokens_t->i++;
        int counter_pow = 0;
        while (str[tokens_t->i] > 47 && str[tokens_t->i] < 58) {
          tokens_t->temp_res =
              tokens_t->temp_res * 10 + (str[tokens_t->i++] - '0');
          counter_pow++;
        }
        tokens_t->temp_res /= pow(10, (long double)counter_pow);
        break;
      }
    }
    tokens_t->values = 1;
    output = get_double(tokens_t);
    if (str[tokens_t->i] == 'e') {
      tokens_t->i = tokens_t->i + 2;
      int i = atoi(&str[tokens_t->i]);
      if (str[tokens_t->i - 1] == '-') {
        for (; i > 0; i--) tokens_t->temp_res = tokens_t->temp_res / 10;
      } else if (str[tokens_t->i - 1] == '+') {
        for (; i > 0; i--) tokens_t->temp_res = tokens_t->temp_res * 10;
      } else
        output = FAILURE;
      while (str[tokens_t->i] > 47 && str[tokens_t->i] < 58) tokens_t->i++;
    }
  }

  return output;
}

double binary_calc(double *value_1, double *value_2, char operator) {
  double result = *value_1;

  switch (operator) {
    case '+':
      result = *value_1 + *value_2;
      break;
    case '-':
      result = *value_1 - *value_2;
      break;
    case '*':
      result = *value_1 * *value_2;
      break;
    case '/':
      result = *value_1 / *value_2;
      break;
    case '%':
      result = fmodl(*value_1, *value_2);
      break;
    case '^':
      result = powl(*value_1, *value_2);
      break;
  }

  return result;
}

int8_t calc_parsing(char *str, struct tokens *tokens_t, node_t *stack) {
  int8_t output = SUCCESS;
  node_t *temp_stack = NULL;

  while (str[tokens_t->i] != '\0' && !output) {
    output = a_to_double(str, tokens_t);
    if (str[tokens_t->i] == '(' && !output) {
      tokens_t->i++;
      tokens_t->brackets++;
      tokens_t = save_operator(tokens_t);
      if (tokens_t->symbols || tokens_t->values) output = FAILURE;
      output = output ? output : calc_parsing(str, tokens_t, temp_stack);
    } else if (str[tokens_t->i] == ')' && !output) {
      tokens_t->i++;
      if (temp_stack == NULL && check_unary(tokens_t)) output = FAILURE;
      output = output ? output : smart_push(&temp_stack, tokens_t);
      output = output ? output : run_calc(&temp_stack, tokens_t);
      output = output ? output : get_double(take_operator(tokens_t));
      tokens_t->values = 1;
      tokens_t->brackets--;
      break;
    }
    if (!output) get_operator(tokens_t, str);
    if ((str[tokens_t->i] == '\0' || str[tokens_t->i] == ')') &&
        tokens_t->symbols)
      output = FAILURE;
    if (!output && tokens_t->values &&
        (tokens_t->symbols || !tokens_t->operator_before_brackets))
      output = smart_push(tokens_t->brackets ? &temp_stack : &stack, tokens_t);
    if (tokens_t->symbols &&
        strchr("+-*/^%cstCSTqLle", str[tokens_t->i]) != NULL)
      output = FAILURE;
  }
  if (!output && !tokens_t->brackets) output = run_calc(&stack, tokens_t);
  remove_stack(&temp_stack);
  return output;
}

int8_t check_unary(struct tokens *tokens_t) {
  return tokens_t->temp_symbols == 'c' || tokens_t->temp_symbols == 's' ||
         tokens_t->temp_symbols == 't' || tokens_t->temp_symbols == 'C' ||
         tokens_t->temp_symbols == 'S' || tokens_t->temp_symbols == 'T' ||
         tokens_t->temp_symbols == 'q' || tokens_t->temp_symbols == 'L' ||
         tokens_t->temp_symbols == 'l';
}

int8_t get_double(struct tokens *tokens_t) {
  int8_t output = SUCCESS;

  if (tokens_t->values) {
    if (tokens_t->temp_symbols == '+' || tokens_t->temp_symbols == '-') {
      if (tokens_t->temp_symbols == '-') tokens_t->temp_res *= -1;
      tokens_t->temp_symbols = 0;
      tokens_t->symbols = 0;
    } else if (tokens_t->symbols &&
               (tokens_t->temp_symbols == '^' ||
                get_priority(tokens_t->temp_symbols) < 3)) {
      output = FAILURE;
    } else if (tokens_t->symbols) {
      tokens_t->temp_res =
          unary_calc(&tokens_t->temp_res, tokens_t->temp_symbols);
      tokens_t->temp_symbols = 0;
      tokens_t->symbols = 0;
    }
  }

  return output;
}

void get_operator(struct tokens *tokens_t, char *str) {
  if (!tokens_t->symbols && str[tokens_t->i] &&
      (str[tokens_t->i] < '0' || str[tokens_t->i] > '9') &&
      (str[tokens_t->i] != '(' && str[tokens_t->i] != ')')) {
    if (str[tokens_t->i] != 'e') tokens_t->temp_symbols = str[tokens_t->i++];
    while (str[tokens_t->i] == '_') tokens_t->i++;
    tokens_t->symbols = 1;
  }
}

int8_t get_priority(char operator) {
  int8_t result = 0;

  switch (operator) {
    case '+':
    case '-':
      result = 1;
      break;
    case '*':
    case '/':
    case '%':
      result = 2;
      break;
    case '^':
    case 'c':
    case 's':
    case 't':
    case 'C':
    case 'S':
    case 'T':
    case 'q':
    case 'L':
    case 'l':
      result = 3;
      break;
  }

  return result;
}

int8_t pop(node_t **stack, double *value, char *operator) {
  int8_t output = FAILURE;
  node_t *temp = NULL;

  if (*stack != NULL) {
    temp = *stack;
    *value = temp->value;
    *operator= temp->symbol;
    *stack = (*stack)->next;
    output = SUCCESS;
    free(temp);
  }

  return output;
}

int8_t push(node_t **stack, double value, char operator) {
  int8_t output = FAILURE;
  node_t *temp = calloc(1, sizeof(node_t));

  if (temp != NULL) {
    temp->value = value;
    temp->symbol = operator;
    temp->next = *stack;
    *stack = temp;
    output = SUCCESS;
  }

  return output;
}

void remove_stack(node_t **stack) {
  node_t *temp = NULL;

  while (*stack != NULL) {
    temp = *stack;
    *stack = (*stack)->next;
    free(temp);
  }
}

void remove_tokens_t(struct tokens *tokens_t) {
  operators_t *temp = NULL;

  while (tokens_t->operator_before_brackets != NULL) {
    temp = tokens_t->operator_before_brackets;
    tokens_t->operator_before_brackets = temp->next;
    free(temp);
  }
  memset(tokens_t, 0, sizeof(*tokens_t));
}

int8_t run_calc(node_t **stacks, struct tokens *tokens_t) {
  int8_t output = SUCCESS;
  node_t *temp_stack = NULL;
  output = stack_rewind(stacks, tokens_t);

  if (!output && !pop(stacks, &tokens_t->temp_res, &tokens_t->temp_symbols)) {
    char temp_ch = tokens_t->temp_symbols;
    double temp_d = tokens_t->temp_res;
    while (!output &&
           !pop(stacks, &tokens_t->temp_res, &tokens_t->temp_symbols)) {
      if (temp_stack ||
          (get_priority(tokens_t->temp_symbols) > get_priority(temp_ch))) {
        if (!temp_stack) {
          output =
              push(&temp_stack, tokens_t->temp_res, tokens_t->temp_symbols);
        } else if (get_priority(tokens_t->temp_symbols) <
                   get_priority(temp_stack->symbol)) {
          temp_stack->value = binary_calc(
              &temp_stack->value, &tokens_t->temp_res, temp_stack->symbol);
          temp_d = tokens_t->temp_res =
              binary_calc(&temp_d, &temp_stack->value, temp_ch);
          temp_ch = tokens_t->temp_symbols;
          remove_stack(&temp_stack);
        } else {
          temp_stack->value = binary_calc(
              &temp_stack->value, &tokens_t->temp_res, temp_stack->symbol);
          temp_stack->symbol = tokens_t->temp_symbols;
        }
      } else {
        temp_d = tokens_t->temp_res =
            binary_calc(&temp_d, &tokens_t->temp_res, temp_ch);
        temp_ch = tokens_t->temp_symbols;
      }
    }
  }
  remove_stack(&temp_stack);
  return output;
}

struct tokens *save_operator(struct tokens *tokens_t) {
  if (check_unary(tokens_t) ||
      (tokens_t->temp_symbols == '+' || tokens_t->temp_symbols == '-')) {
    operators_t *temp = calloc(1, sizeof(operators_t));
    if (temp != NULL) {
      temp->symbol = tokens_t->temp_symbols;
      temp->brackets = tokens_t->brackets;
      temp->next = tokens_t->operator_before_brackets;
      tokens_t->operator_before_brackets = temp;
      tokens_t->temp_symbols = 0;
      tokens_t->symbols = 0;
    }
  }

  return tokens_t;
}

int8_t smart_push(node_t **stacks, struct tokens *tokens_t) {
  int8_t output = SUCCESS;
  if (check_unary(tokens_t)) output = FAILURE;
  if (*stacks != NULL && !output) {
    if ((*stacks)->symbol == '^') {
      (*stacks)->value = binary_calc(&(*stacks)->value, &tokens_t->temp_res,
                                     (*stacks)->symbol);
      (*stacks)->symbol = tokens_t->temp_symbols;
    } else
      output = push(stacks, tokens_t->temp_res, tokens_t->temp_symbols);
  } else if (!output)
    output = push(stacks, tokens_t->temp_res, tokens_t->temp_symbols);

  tokens_t->values = 0;
  tokens_t->symbols = 0;
  tokens_t->temp_res = 0;
  tokens_t->temp_symbols = 0;
  return output;
}

int8_t stack_rewind(node_t **stacks, struct tokens *tokens_t) {
  int8_t output = SUCCESS;
  node_t *temp_stack = NULL;

  while (!pop(stacks, &tokens_t->temp_res, &tokens_t->temp_symbols) && !output)
    output = push(&temp_stack, tokens_t->temp_res, tokens_t->temp_symbols);
  remove_stack(stacks);
  if (temp_stack != NULL) *stacks = temp_stack;

  return output;
}

struct tokens *take_operator(struct tokens *tokens_t) {
  operators_t *temp = NULL;
  if (tokens_t->operator_before_brackets != NULL) {
    if (tokens_t->operator_before_brackets->brackets == tokens_t->brackets) {
      temp = tokens_t->operator_before_brackets;
      tokens_t->temp_symbols = temp->symbol;
      tokens_t->operator_before_brackets = temp->next;
      tokens_t->symbols = 1;
      tokens_t->values = 1;
      free(temp);
    } else
      tokens_t->temp_symbols = 0;
  }

  return tokens_t;
}

double unary_calc(double *value, char operator) {
  double result = 0.0;

  switch (operator) {
    case 'c':
      result = cosl(*value);
      break;
    case 's':
      result = sinl(*value);
      break;
    case 't':
      result = tanl(*value);
      break;
    case 'C':
      result = acosl(*value);
      break;
    case 'S':
      result = asinl(*value);
      break;
    case 'T':
      result = atanl(*value);
      break;
    case 'q':
      result = sqrtl(*value);
      break;
    case 'l':
      result = logl(*value);
      break;
    case 'L':
      result = log10l(*value);
      break;
  }

  return result;
}
