# Dungeons & Dragons Assistant

A library and tool suite written in Go for tracking things like character status, items, and equipment for D&D 5e.

## DnD CLI

The assistant is a command line interface tool used to interact with the data in the database.

## Library Design

There are layers of abstraction to isolate the behaviors of library.

Entity
: The item or object to be tracked. Ex. A character or Weapon

Service
: The manipulation methods available to the entity

Repository
: Implements a service interface and contains logic to interact with a specific database