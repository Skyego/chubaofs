### Command examples

```example bash
使用方式：
volfile：每一行一个volname
nohup ./mefsck-cn check both --master "127.0.0.1:17010" -f "vol-cn.txt" -b 30 --mport "9092" > cn.out &

./fsck check inode ----master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck check dentry --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck check both --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck check both --vol "<volName>" --inode-list "inodes.txt" --dentry-list "dens.txt"
./fsck clean evict --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck clean inode --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck clean inode --vol "<volName>" --inode-list "inodes.txt" --dentry-list "dens.txt"
./fsck clean dentry --master "127.0.0.1:17010" --vol "<volName>" --mport "17220"
./fsck clean dentry --vol "<volName>" --inode-list "inodes.txt" --dentry-list "dens.txt"
```
