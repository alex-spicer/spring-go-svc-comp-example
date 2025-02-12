package com.example.product.repository;

import com.example.product.model.Product;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicLong;

@Component
public class InMemoryRepository {
    private final Map<Long, Product> products = new ConcurrentHashMap<>();
    private final AtomicLong idGenerator = new AtomicLong(1);

    public List<Product> getAllProducts() {
        return List.copyOf(products.values());
    }

    public Optional<Product> getProduct(Long id) {
        return Optional.ofNullable(products.get(id));
    }

    public Product createProduct(Product product) {
        Long id = idGenerator.getAndIncrement();
        Product newProduct = product.withId(id);
        products.put(id, newProduct);
        return newProduct;
    }

    public Optional<Product> updateProduct(Long id, Product product) {
        return Optional.ofNullable(products.replace(id, product.withId(id)));
    }

    public boolean deleteProduct(Long id) {
        return products.remove(id) != null;
    }
}
