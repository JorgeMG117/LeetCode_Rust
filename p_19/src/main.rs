struct Solution {}

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>
}

impl ListNode {
//    #[inline]
//    fn new(val: i32) -> Self {
//         ListNode {
//             next: None,
//             val
//         }
//     }
}

impl Solution {
    fn remove_nth_from_end_req(head: &mut Option<Box<ListNode>>, n: i32) -> i32 {
        if let Some(node) = head {
            let val = Self::remove_nth_from_end_req(&mut node.next, n);
            if val == n {
                let nodo_val = node.val;
                println!("Eliminar nodo {nodo_val}");

                *head = head.as_mut().unwrap().next.take();
            }
            return val + 1;
        } else {
            return 1;
        }
    }

    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        let mut head = head;
        Self::remove_nth_from_end_req(&mut head, n);
        return head;
    }
}

// 19. Remove Nth Node From End of List
fn main() {
    //let mut head = Some(Box::new(ListNode { val: 1, next: None }));
    //let mut head = Some(Box::new(ListNode { val: 1, next: Some(Box::new(ListNode { val: 2, next: None })) }));
    //let mut head = Some(Box::new(ListNode { val: 1, next: Some(Box::new(ListNode { val: 2, next: Some(Box::new(ListNode { val: 3, next: None })) })) }));
    let head = Some(Box::new(ListNode { val: 1, next: Some(Box::new(ListNode { val: 2, next: Some(Box::new(ListNode { val: 3, next: Some(Box::new(ListNode { val: 4, next: None })) })) })) }));
    let n = 2;
    
    let sol = Solution::remove_nth_from_end(head, n);

    //Print the elements of the list
    let mut curr = sol;
    while let Some(node) = curr {
        let value = node.val;
        println!("{value}");
        curr = node.next;
    }
}
