function tryingRecursion(i) {
    if (i == 0) {
        return 0; 
    }
    console.log(i)
    return tryingRecursion(i - 1);
}

//console.log(tryingRecursion(20))


function countDown(n) {
    for (let i = n; i >=0; i--) {
        console.log(i)
    }

}

countDown(20)