# timecard
A command line tool to track time spent on projects

## Design

This tool stores data about what projects are being worked on in
plain-text files.  The motivation of this decision is to allow a text
editor to be used to inspect and modify the database for debugging or
error-correction purposes.  The format of the file is one entry per
line, each line consisting of an RFC 3339 date and time and the
identifier of the project.  These files are stored, one per day, under
the directory `~/.config/timecard/data`.

There is also a file which gives descriptions of the projects.  This
file also has one entry per line; the first word is the project
identifier, and the rest of the line is the description.  This file is
stored as `~/.config/timecard/projects`.

## Motivation

I have to report my time in hours charged to different projects
depending on what I work on during the day.  This tool assists with
reporting the time accurately.  This is also an opportunity to use
golang.
