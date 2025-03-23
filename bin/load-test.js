import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    stages: [
        { duration: '30s', target: 5 },    // Ramp up to 5 users
        { duration: '1m', target: 5 },     // Stay at 5 users
        { duration: '20s', target: 10 },   // Ramp up to 10 users
        { duration: '1m', target: 10 },    // Stay at 10 users
        { duration: '20s', target: 0 },    // Scale down to 0 users
    ],
    thresholds: {
        http_req_duration: ['p(95)<2000'], // 95% of requests must complete below 2s
        http_req_failed: ['rate<0.1'],     // Less than 10% can fail
    },
};

const BASE_URL = 'http://localhost:8080/api/v1';

// Generates a random UUID v4
function uuidv4() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        const r = Math.random() * 16 | 0;
        const v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

export default function () {
    const userId = uuidv4();
    const authToken = `Bearer ${randomString(32)}`;

    // Add trace context headers
    const headers = {
        'Authorization': authToken,
        'traceparent': `00-${randomString(32)}-${randomString(16)}-01`,
        'Content-Type': 'application/json',
    };

    // Test user endpoint with trace context
    const userResponse = http.get(`${BASE_URL}/users/${userId}`, {
        headers: headers,
        tags: { name: 'GetUserByID' },
    });

    // Check response and add custom span attributes via tags
    check(userResponse, {
        'status is 401 or 422': (r) => r.status === 401 || r.status === 422,
        'response time OK': (r) => r.timings.duration < 2000,
    });

    // Add some variation in the test pattern
    if (Math.random() < 0.3) {
        // Simulate slow requests occasionally
        sleep(2);
    } else {
        sleep(1);
    }

    // Test invalid UUID to generate error traces
    if (Math.random() < 0.1) {
        const invalidResponse = http.get(`${BASE_URL}/users/invalid-uuid`, {
            headers: headers,
            tags: { name: 'GetUserByIDInvalid' },
        });

        check(invalidResponse, {
            'invalid uuid returns 400': (r) => r.status === 400,
        });
    }
}