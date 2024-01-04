const val = {
    a: 2,
    b: 3,
    test: [1, 2, 4, 5],
    tester: [10, 9, 0, 54, "test"],
    val2: {
        b: 3,
        c: 5,
        ayoAges: [10, 20, 50, "Aduwo"],
        isAyCorrect: true,
   }
};

let numberToString = (obj) => Object.keys(obj).forEach((key) => {
    const value = obj[key];

    if (!Array.isArray(value) && typeof value == "number") { 
        obj[key] = value.toString()
        }
        else if (Array.isArray(value)) {

            const strArrayVal = value.map((val) => {
                if (typeof val != "number") {
                    return val;
                }
                return val.toString()
                
            })
       obj[key] = strArrayVal
    }
    else if (typeof value === "object") {
        numberToString(value)
    }

})

numberToString(val)
console.log(val)


function numberToString2(obj) {
    for (const key in obj) {
        if (!Array.isArray(obj[key]) && typeof obj[key] === "number") { 
            obj[key] = obj[key].toString();
        } else if (Array.isArray(obj[key])) {
            const strArrayVal = obj[key].map((val) => {
                if (typeof val !== "number") {
                    return val;
                }
                return val.toString();
            });
            obj[key] = strArrayVal;
        } else if (typeof obj[key] === "object") {
            numberToString2(obj[key]); 
        }
    }
}

