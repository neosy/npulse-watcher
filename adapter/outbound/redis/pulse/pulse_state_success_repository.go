package rpulse

import (
	"context"
	"fmt"
)

// AddSuccess saves ip as success.
func (r *PulseStateRepository) AddSuccess(ctx context.Context, ip string) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "success")

	err := r.client.SAdd(ctx, key, ip).Err()
	if err != nil {
		return fmt.Errorf("failed to add value in Redis: %v", err)
	}

	return nil
}

// RemoveSuccess unmarks the given IP as successful.
func (r *PulseStateRepository) RemoveSuccess(ctx context.Context, ip string) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "success")

	err := r.client.SRem(ctx, key, ip)
	if err != nil {
		return fmt.Errorf("failed to remove value in Redis: %v", err)
	}

	return nil
}

// RemoveAllSuccess clears all successful IPs.
func (r *PulseStateRepository) RemoveAllSuccess(ctx context.Context) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "success")

	err := r.client.Del(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to remove key in Redis: %v", err)
	}

	return nil
}

// GetAllSuccessful returns all successful ips
func (r *PulseStateRepository) GetAllSuccessful(ctx context.Context) ([]string, error) {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "success")

	ips, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get in Redis: %v", err)
	}

	return ips, nil
}
