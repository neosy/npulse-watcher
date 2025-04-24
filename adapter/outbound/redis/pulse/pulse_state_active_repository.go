package rpulse

import (
	"context"
	"fmt"
)

// AddActiveIP
func (r *PulseStateRepository) AddActiveIP(ctx context.Context, ip string) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "active_ips")

	err := r.client.SAdd(ctx, key, ip).Err()
	if err != nil {
		return fmt.Errorf("failed to add value in Redis: %v", err)
	}

	return nil
}

// RemoveActiveIP
func (r *PulseStateRepository) RemoveActiveIP(ctx context.Context, ip string) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "active_ips")

	err := r.client.SRem(ctx, key, ip)
	if err != nil {
		return fmt.Errorf("failed to remove value in Redis: %v", err)
	}

	return nil
}

// RemoveAllActiveIPs
func (r *PulseStateRepository) RemoveAllActiveIPs(ctx context.Context) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "active_ips")

	err := r.client.Del(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to remove key in Redis: %v", err)
	}

	return nil
}

// GetAllActiveIPs
func (r *PulseStateRepository) GetAllActiveIPs(ctx context.Context) ([]string, error) {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "active_ips")

	ips, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get in Redis: %v", err)
	}

	return ips, nil
}
