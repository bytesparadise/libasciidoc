# When you run make VERBOSE=1 (the default), executed commands will be printed
# before executed. If you run make VERBOSE=2 verbose flags are turned on and
# quiet flags are turned off for various commands. Use V_FLAG in places where
# you can toggle on/off verbosity using -v. Use Q_FLAG in places where you can
# toggle on/off quiet mode using -q. Use S_FLAG where you want to toggle on/off
# silence mode using -s...
VERBOSE ?= 1
Q = @
Q_FLAG = -q
QUIET_FLAG = --quiet
V_FLAG =
S_FLAG = -s
X_FLAG =
ifeq ($(VERBOSE),1)
       Q =
endif
ifeq ($(VERBOSE),2)
       Q =
       Q_FLAG =
       QUIET_FLAG =
       S_FLAG =
       V_FLAG = -v
       X_FLAG = -x
endif