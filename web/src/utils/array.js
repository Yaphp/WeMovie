/**
 * 从数组中随机抽取几个元素
 * @param array
 * @param n
 * @returns {*}
 */
function choice(array, n = 1) {
    var counter = array.length, temp, index;

    // While there are elements in the array
    while (counter--) {
        // Pick a random index
        index = (Math.random() * counter) | 0;

        // And swap the last element with it
        temp = array[counter];
        array[counter] = array[index];
        array[index] = temp;
    }
    array.length = n
    return array;
}

/**
 * 从数组中随机抽取1个元素
 * @param array
 * @param n
 * @returns {*}
 */
function choiceOne(array, n = 1) {
    if (array.length === 0) {
        throw new Error('Cannot choose from an empty array')
    }
    var counter = array.length, temp, index;

    // While there are elements in the array
    while (counter--) {
        // Pick a random index
        index = (Math.random() * counter) | 0;

        // And swap the last element with it
        temp = array[counter];
        array[counter] = array[index];
        array[index] = temp;
    }
    array.length = n
    return array[0];
}

export {choice, choiceOne}