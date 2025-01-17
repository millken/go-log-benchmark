##
# gnuplot script to generate a performance graphic.
#
# it expects the following parameters:
#
# file_path - path to the file from which the data will be read
# graphic_file_name - the graphic file name to be saved 
# y_label - the desired label for y axis
# y_range_min - minimum range for values in y axis
# y_range_max - maximum range for values in y axis
# column_1 - the first column to be used in plot command
# column_2 - the second column to be used in plot command
#
# Author: Tiago Melo (tiagoharris@gmail.com)
##

# graphic will be saved as 800x600 png image file
set terminal png

# allows grid lines to be drawn on the plot
set grid

# setting the graphic file name to be saved
set output graphic_file_name

# the graphic's main title
set title "performance comparison"

# since the input file is a CSV file, we need to tell gnuplot that data fields are separated by comma
set datafile separator ","

# disable key box
set key off

# label for y axis
set ylabel y_label

# range for values in y axis
set yrange[y_range_min:y_range_max]

# to avoid displaying large numbers in exponential format
set format y "%.0f"

# vertical label for x values 
set xtics rotate

# set boxplots
set style fill solid
set boxwidth 0.5

# plot graphic for each line of input file
plot for [i=0:*] file_path every ::i::i using column_1:column_2:xtic(2) with boxes