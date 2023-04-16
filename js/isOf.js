var checkIfInstanceOf = function(obj, classFunction) {
    // 1
    // if (obj == null || classFunction == null) {
    //     return false;
    // }
    // if (obj instanceof classFunction) {
    //     return true;
    // } else if (obj.prototype instanceof classFunction) {
    //     return true;
    // } else if (typeof(obj) === classFunction.name.toLocaleLowerCase()) {
    //     return true;
    // }
    // return false;

    // 2
    // if (obj === null || obj === undefined) {
    //     return false;
    //   }
    //   if (obj.constructor === classFunction) {
    //     return true;
    //   }
    //   let prototype = Object.getPrototypeOf(obj);
    //   while (prototype !== null) {
    //     if (prototype.constructor === classFunction) {
    //       return true;
    //     }
    //     prototype = Object.getPrototypeOf(prototype);
    //   }
    //   return false;

    if(obj === null || obj === undefined || (typeof classFunction !== 'function')){
        return false
    }
    return Object(obj) instanceof classFunction
};

console.log(checkIfInstanceOf(null, null));
console.log(checkIfInstanceOf(new Date(), Date));
console.log(checkIfInstanceOf([], []));