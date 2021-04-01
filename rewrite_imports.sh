find . \( -type d -name .git -prune \) -o -type f -print0 | \
  xargs -0 sed -i "s;github.com/aristanetworks/hdfs;github.com/aristanetworks/hdfs;g"
