# DIRS - Decentralized Information Retrieval System
A project for my diploma in Saint Petersburg's State University

# Quickly writing architecture, so i will not forget

## Terminology

### Request VS Task
Request is foreign piece of information.
Task is a formated request

## Components

### Listener
Interpretes requests from peers and sends them to the *thinker*

### Broadcaster
Interpretes tasks received from the *thinker* into requests to another nodes

### Economy
Formulates storage and demand of system and network. Used by *thinker* to formulate new tasks

### Butler
Interpretes requests from user to give to the *thinker*

### Master
A user cli to communicate with *butler*

### Thinker
Component that formulates tasks from input and watches their execution