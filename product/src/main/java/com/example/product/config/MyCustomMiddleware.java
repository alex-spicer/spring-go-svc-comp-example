package com.example.product.config;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

@Component
public class MyCustomMiddleware implements HandlerInterceptor {
    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
            throws Exception {
        if ("GET".equalsIgnoreCase(request.getMethod())) {
            System.out.println("PreHandle - GET Request to: " + request.getRequestURI());
        }
        return true; // Return true to continue processing, false to abort
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler,
                           ModelAndView modelAndView) throws Exception {
        if ("GET".equalsIgnoreCase(request.getMethod())) {
            System.out.println("PostHandle - GET Response status: " + response.getStatus());
        }
    }

}
