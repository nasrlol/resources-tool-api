
# SYNF

A tool that allows developers to monitor devices from anywhere, at any time, however they want.
You can find more information on the website at www.nsrddyn.com. The website is built in React.

In this repository, you can find the full source code.

Everything starts at **synf-sys**,
the source of the data, written in C. This is what pulls the data from the user's device.
A main goal of the program is to minimize resource consumption; at the moment, this has not been tested or explored yet.

The API that supports the application consists of the following routes:

```
GET /api/data/device
GET /api/data/cpu
GET /api/data/gpu
GET /api/data/ram
GET /api/data/disk

```

