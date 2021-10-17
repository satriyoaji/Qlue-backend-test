#!/bin/sh

extension=".txt"
extension_length=`echo ".txt"|wc -c`
path="mig33/inner_folder/"

for i in `ls $path*$extension`;
do
  length=`echo $i|wc -c`

  filename_length=`expr $length - $extension_length`

  filename=`echo $i| cut -c $filename_length`
  filename=$filename.dat

  mv $i $path$filename
done