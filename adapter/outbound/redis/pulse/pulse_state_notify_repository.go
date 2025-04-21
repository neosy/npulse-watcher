package rpulse

import (
	"context"
	"fmt"
)

// AddNotifyIP
func (r *PulseStateRepository) AddNotifyIP(ctx context.Context, ip string) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "notify_ips")

	err := r.client.SAdd(ctx, key, ip).Err()
	if err != nil {
		return fmt.Errorf("failed to add value in Redis: %v", err)
	}

	return nil
}

// RemoveNotifyIP
func (r *PulseStateRepository) RemoveNotifyIP(ctx context.Context, ip string) error {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "notify_ips")

	err := r.client.SRem(ctx, key, ip)
	if err != nil {
		return fmt.Errorf("failed to remove value in Redis: %v", err)
	}

	return nil
}

// GetAllActiveIPs
func (r *PulseStateRepository) GetAllNotifyIPs(ctx context.Context) ([]string, error) {
	key := fmt.Sprintf("%s:%s", r.getBaseKey(), "notify_ips")

	ips, err := r.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get in Redis: %v", err)
	}

	return ips, nil
}
