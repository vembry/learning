import model.LRUCache;

public class App {
    public static void main(String[] args) throws Exception {
        LRUCache cache = new LRUCache(7);

	    cache.Set("key1", "asd");
	    cache.Set("key2", "asd");
	    cache.Set("key1", "asd");
	    cache.Set("key3", "asd");
	    cache.Set("key4", "asd");
	    cache.Set("key5", "asd");
	    cache.Get("key2");
	    cache.Set("key6", "asd");
	    cache.Set("key7", "asd");
	    cache.Get("key2");
	    cache.Set("key8", "asd");
	    cache.Set("key9", "asd");
	    cache.Set("key10", "asd");
	    cache.Set("key1", "asd");

        cache.print();
    }
}
