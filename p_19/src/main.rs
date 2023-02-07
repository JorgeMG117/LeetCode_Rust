struct Solution {}

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>
}

impl ListNode {
    #[inline]
   fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val
        }
    }
}

impl Solution {
    fn remove_nth_from_end_req(head: Option<Box<ListNode>>, n: i32) -> i32 {
        //si es None
            //devolvemos 1
        //sino
            //val = f(next)
            //si val = n //hay que eliminar este nodo
                //apuntamos ha este nodo
                //con el que nos han llamado es igual a esteNodo.next
            //sino return val+1
        if let Some(node) = head {
            let val = Self::remove_nth_from_end_req(node.next, n);
            if val == n {
                let nodo_val = node.val;
                println!("Eliminar nodo {nodo_val}");

                let aux = node;
                node.next = aux.next;
            }
            return val + 1;
        } else {
            return 1;
        }
    }

    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        //let list = ListNode::new(3);
        //let mut value = list.val;
        //println!("{value}");
        
        /*if let Some(v) = head {
            let value = v.val;
            println!("{value}");
        }*/


        /* 
        //add additional node to handle case of head removal
        let mut dummy = Some(Box::new(ListNode {
            val: -1,
            next: head,
        }));

        //calculate length of the list
        let mut len = 0;
        let mut temp = dummy.as_ref();
        while temp.is_some() {
            temp = temp.unwrap().next.as_ref();
            len += 1;
        }

        //calculate the index to be removed
        let index = len - n - 1;

        //get to the node before the index
        let mut dummy_mut = dummy.as_mut().unwrap();
        for _ in 0..index {
            dummy_mut = dummy_mut.next.as_mut().unwrap();
        }

        //remove the node
        let next_node = dummy_mut.next.take();
        if let Some(node_to_remove) = next_node {
            dummy_mut.next = node_to_remove.next;
        }

        return dummy.unwrap().next;
        */


        Self::remove_nth_from_end_req(head, 32);
        return head;

    }
}

// 19. Remove Nth Node From End of List
fn main() {
    println!("Hello, world!");
}
