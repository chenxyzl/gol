
svn cleanup D:\\Work\\sanguo\\svn\\configCsv\\csv
svn up D:\\Work\\sanguo\\svn\\configCsv\\csv
cd ../../data
cp -r D:\\Work\\sanguo\\svn\\configCsv\\csv\\. ./

filelist=`ls .`
for file in $filelist
do 
    dos2unix $file
done

read -p "按任意键继续" -n 1 -r
