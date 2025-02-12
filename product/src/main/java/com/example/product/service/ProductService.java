package com.example.product.service;

import com.example.product.model.Product;
import com.example.product.repository.InMemoryRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ProductService {
    private final InMemoryRepository repository;

    public ProductService(InMemoryRepository repository) {
        this.repository = repository;
    }

    public List<Product> getAllProducts() {
        return repository.getAllProducts();
    }

    public Optional<Product> getProduct(Long id) {
        return repository.getProduct(id);
    }

    public Product createProduct(Product product) {
        return repository.createProduct(product);
    }

    public Optional<Product> updateProduct(Long id, Product product) {
        return repository.updateProduct(id, product);
    }

    public boolean deleteProduct(Long id) {
        return repository.deleteProduct(id);
    }
}