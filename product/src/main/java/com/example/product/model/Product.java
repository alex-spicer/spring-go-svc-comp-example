package com.example.product.model;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Positive;

public record Product(
        Long id,

        @NotBlank(message = "Product name is required")
        String name,

        @Positive(message = "Price must be positive")
        double price,

        String description
) {
    public Product withId(Long id) {
        return new Product(id, name(), price(), description());
    }
}