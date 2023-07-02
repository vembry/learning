package model;

public class Node {
    public String key;
    public String val;
    public Node prev;
    public Node next;
    public Node(String key, String val){
        this.key = key;
        this.val = val;
    }
}
