
function printArrayResult(arr)
{
    if (arr.length !== 1) {

        for (let i = 0; i < arr.length; i++) {
            document.write( arr[i] + " ");
        }
        document.write("<br>");
    }
}

function findNumbers(arr, i, n)
{
    if (n === 0)
        printArrayResult(arr);

    for (let j = i; j <= n; j++) {

        arr.push(j);

        findNumbers(arr, j, n - j);

        arr.pop();
    }
}

var n = 4;
var arr = [];

findNumbers(arr, 1, n);

