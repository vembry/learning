package model;

import java.util.HashMap;

public class LRUCache {
    public Integer cap;
    public Node head;
    public Node tail;
    
    public HashMap<String, Node> cacheMap;

    public LRUCache(Integer cap) {
        this.cap = cap;

        this.head = new Node("head", "head");
        this.tail = new Node("tail", "tail");
        
        // initiatlize head
        this.head.prev = this.tail;
        this.head.next = null;

        // initialize tail
        this.tail.prev = null;
        this.tail.next = this.head;

        this.cacheMap = new HashMap<>();
    }

    public void Set(String key, String val) {
        System.out.printf("inserting key=%s, val=%s", key, val);
        System.out.println("");

        Node newNode = new Node(key, val);

        if(this.cacheMap.containsKey(key)){
            // if key already exists
            newNode = this.takeOutNode(this.cacheMap.get(key));
        } else if (this.cacheMap.size() == this.cap) {
            // if cache reached cap threshold
            Node leastUsedNode = this.takeOutNode(this.tail.next);
            this.cacheMap.remove(leastUsedNode.key);
        }

        this.cacheMap.put(key, newNode);
        this.setNodeOnTop(newNode);
    }

    private void setNodeOnTop(Node node) {
        node.next = this.head;
        node.prev = this.head.prev;
        node.prev.next = node;
        this.head.prev = node;
    }

    private Node takeOutNode(Node node){
        node.next.prev = node.prev;
        node.prev.next = node.next;
        node.next = node.prev = null;
        return node;
    }

    public String Get(String key) {
        System.out.printf("getting key=%s", key);
        System.out.println("");
        return "";
    }

    public void print() {
        System.out.println("printing-----------------------------------");

        Node curr = this.tail.next;
        while (curr.next != null) {
            System.out.printf("key=%s, val=%s", curr.key, curr.val);
            System.out.println("");
            curr = curr.next;
        }
    }
}
