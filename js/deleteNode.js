class ListNode {
    constructor(val) {
        this.val = val;
        this.next = null;
    }
}
 
/**
 * @param {ListNode} node
 * @return {void} Do not return anything, modify node in-place instead.
 */
 var deleteNode = function(node) {
    node.val = node.next.val;
    node.next = node.next.next
};

/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    let next = head
    let prev = head
    
    for (let index = 0; index < n; index++) {
        next = next.next;
    }
    
    if(next === null) return head.next;

    while(next.next !== null) {
        prev = prev.next
        next = next.next
    }

    prev.next = prev.next.next

    return head
};