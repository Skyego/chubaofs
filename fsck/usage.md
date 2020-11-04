### Command examples

```example bash
使用方式：
volfile：每一行一个volname
nohup ./mefsck-cn check both --master "127.0.0.1:17010" -f "vol-cn.txt" -b 20 --mport "9092" > cn.out &

./fsck check inode ----master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck check dentry --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck check both --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck check both --vol "<volName>" --inode-list "inodes.txt" --dentry-list "dens.txt"
./fsck clean evict --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck clean inode --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck clean inode --vol "<volName>" --inode-list "inodes.txt" --dentry-list "dens.txt"
./fsck clean dentry --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck clean dentry --vol "<volName>" --inode-list "inodes.txt" --dentry-list "dens.txt"

结果数据说明：
{"VolName":"hadoop_test_vol10",
"TotalFileSize":2507723960,
"ObsoleteTotalFileSize":2788,
"SafeCleanSize":2788},

第一个数据是通过mp统计的总共的数据大小
第二个是孤儿inode占用的空间大小

第一个包含第二个，通过getCluster得到的数据，包含并且>=第一个

```
