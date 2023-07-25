let array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
// 配列の中身だけfor文を繰り返す。i = 0 ++を使いたい
for (let i = 0; i < array.length; i++) {
  if(i == 8) {
    console.log(array[i]);
    break;
  }
}
