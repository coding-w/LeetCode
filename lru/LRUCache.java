package lru;

import java.util.HashMap;

public class LRUCache<K, V> {

    // Node 双向链表
    private class Node {
        // 键
        private K key;
        // 值
        private V value;
        // 前置节点
        private Node prev;
        // 后置节点
        private Node next;

        public Node(K key, V value) {
            this.key = key;
            this.value = value;
        }
        public Node(){}
    }

    // 最大缓存数量
    private final Integer maxSize;

    // 头节点
    private final Node head;

    // 尾节点
    private final Node tail;

    // Hash
    private final HashMap<K, Node> map;

    // 构造方法
    public LRUCache(int maxSize) {
        if (maxSize <= 0) {
            throw new IllegalArgumentException("Cache size must be positive");
        }
        this.maxSize = maxSize;
        head = new Node();
        tail = new Node();
        head.next = tail;
        tail.prev = head;
        map = new HashMap<K, Node>();
    }

    // 获取键值
    public V get(K key) {
        Node node = map.get(key);
        if (node == null) {
            return null;
        }
        nodeToHead(node);
        return node.value;
    }

    // 添加元素
    public void put(K key, V value) {
        Node node = map.get(key);
        if (node != null) {
            node.value = value;
            nodeToHead(node);
        } else {
            node = new Node(key, value);
            if (map.size() >= maxSize) {
                map.remove(tail.prev.key);
                removeNode(tail.prev);
            }
            map.put(key, node);
            addHeadNode(node);
        }
    }

    // 节点前置
    private void nodeToHead(Node node) {
        removeNode(node);
        addHeadNode(node);
    }

    // 添加头节点
    private void addHeadNode(Node node) {
        node.next = head.next;
        node.prev = head;
        head.next.prev = node;
        head.next = node;
    }

    // 删除节点
    private void removeNode(Node node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }

    // 打印
    private void print() {
        Node item = head.next;
        while (item.next != null) {
            System.out.println(item.key + " : " + item.value.toString());
            item = item.next;
        }
    }

    public static void main(String[] args) {
        LRUCache<Integer, String> cache = new LRUCache<>(3);
        cache.put(1, "aaa");
        cache.put(2, "bbb");
        cache.put(3, "ccc");
        System.out.println(cache.get(1));
        System.out.println(cache.get(2));
        cache.put(4, "ddd");
        cache.print();
    }


}
