#!/bin/bash

CURRDIR=`dirname "$0"`
BASEDIR=`cd "$CURRDIR"; pwd`
#VERSION="__BUILD_VERSION__"

NAME="__PROJECT_NAME__"

CMD="__PROJECT_NAME__"

if [ "$1" = "-v" ]; then
	echo "$NAME current version: $VERSION"
    exit
fi

if [ "$1" = "-d" ]; then
    shift
    EXECUTEDIR=$1
    shift
else
    EXECUTEDIR=$BASEDIR
fi

if [ ! -d "$EXECUTEDIR" ]; then
    echo "ERROR: $EXECUTEDIR is not a dir"
    exit
fi

if [ ! -d "$EXECUTEDIR"/conf ]; then
    echo "ERROR: could not find $EXECUTEDIR/conf/"
    exit
fi

if [ ! -d "$EXECUTEDIR"/runtime ]; then
    mkdir "$EXECUTEDIR"/runtime
fi

if [ ! -d "$EXECUTEDIR"/runtime/logs ]; then
    mkdir "$EXECUTEDIR"/runtime/logs
fi

cd "$EXECUTEDIR"

PID_FILE="$EXECUTEDIR"/runtime/"$NAME".pid

check_pid() {
    RETVAL=1
    if [ -f $PID_FILE ]; then
        PID=`cat $PID_FILE`
        ls /proc/$PID &> /dev/null
        if [ $? -eq 0 ]; then
            RETVAL=0
        fi
    fi
}

check_running() {
    PID=0
    RETVAL=0
    check_pid
    if [ $RETVAL -eq 0 ]; then
        echo "$CMD is running as $PID, we'll do nothing"
        exit
    fi
}


start() {
    check_running
    echo "starting $CMD ..."
    #"$BASEDIR"/bin/"$CMD" -d "$EXECUTEDIR" 2> "$EXECUTEDIR"/runtime/logs/app.err > "$EXECUTEDIR"/runtime/logs/app.out &
    "$BASEDIR"/bin/"$CMD" -d "$EXECUTEDIR" 2> "$EXECUTEDIR"/runtime/logs/app.err > /dev/null &
    PID=$!
    echo $PID > "$PID_FILE"
	sleep 1
	status
}

stop() {
    check_pid
    if [ $RETVAL -eq 0 ]; then
        echo "$CMD is running as $PID, stopping it..."
        kill -15 $PID
		sleep 1
        echo "done"
    else
        echo "$CMD is not running, do nothing"
    fi

	while true; do
		check_pid
		if [ $RETVAL -eq 0 ]; then
			echo "$CMD is running, waiting it's exit..."
			sleep 1
		else
			echo "$CMD is stopped safely, you can restart it now"
			break
		fi
	done

    if [ -f $PID_FILE ]; then
        rm $PID_FILE
    fi
}

status() {
    check_pid
    if [ $RETVAL -eq 0 ]; then
        echo "$CMD is running as $PID ..."
    else
        echo "$CMD is not running"
    fi
}

RETVAL=0
case "$1" in
    start)
        start $@
        ;;  
    stop)
        stop
        ;;  
    restart)
        stop
        start $@
        ;;  
    status)
        status
        ;;  
    *)  
        #echo "Version: $VERSION"
        echo "Usage: $0 [-d EXECUTION_PATH] {start|stop|restart|status}"
        RETVAL=1
esac
exit $RETVAL

