#ifndef YAGFM_CMARK_BRIDGE_H
#define YAGFM_CMARK_BRIDGE_H
#include <stddef.h>
#include <stdint.h>

typedef struct bridge_document bridge_document;

bridge_document *bridge_parse(const char *source, size_t length);
void bridge_free(bridge_document *document);
char *bridge_render_html(bridge_document *document);
void bridge_free_string(char *value);
void bridge_register_extensions(void);
const char *bridge_version(void);
void *bridge_root(bridge_document *document);

const char *bridge_node_type(void *node);
void *bridge_first_child(void *node);
void *bridge_next(void *node);
void *bridge_parent(void *node);
const char *bridge_node_literal(void *node);
int bridge_set_literal(void *node, const char *literal, size_t length);
int bridge_start_line(void *node);
int bridge_start_column(void *node);
int bridge_end_line(void *node);
int bridge_end_column(void *node);
int bridge_set_extension(void *node, const char *name);
int bridge_set_task_checked(void *node, int checked);
int bridge_get_task_checked(void *node);
#endif
