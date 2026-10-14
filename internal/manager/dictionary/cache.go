package dictionary

import (
	"context"
	"fmt"
	"sync"
	"time"

	rcmodel "github.com/web-rabis/reader-client/model"
)

const cacheTTL = time.Hour

type cacheEntry struct {
	expiresAt time.Time
	count     int64
	data      any
}

// CachingManager — декоратор над IManager, кеширующий ответы справочников
// в памяти на cacheTTL. Справочники меняются редко, поэтому нет смысла
// дёргать reader-client по gRPC на каждый запрос.
type CachingManager struct {
	next IManager

	mu    sync.Mutex
	items map[string]cacheEntry
}

func NewCachingManager(next IManager) *CachingManager {
	return &CachingManager{
		next:  next,
		items: make(map[string]cacheEntry),
	}
}

func cached[T any](c *CachingManager, method string, paging *rcmodel.Paging, fetch func() (int64, T, error)) (int64, T, error) {
	key := cacheKey(method, paging)

	c.mu.Lock()
	if e, ok := c.items[key]; ok && time.Now().Before(e.expiresAt) {
		c.mu.Unlock()
		return e.count, e.data.(T), nil
	}
	c.mu.Unlock()

	count, data, err := fetch()
	if err != nil {
		var zero T
		return 0, zero, err
	}

	c.mu.Lock()
	c.items[key] = cacheEntry{expiresAt: time.Now().Add(cacheTTL), count: count, data: data}
	c.mu.Unlock()

	return count, data, nil
}

func cacheKey(method string, paging *rcmodel.Paging) string {
	if paging == nil {
		return method
	}
	return fmt.Sprintf("%s|%d|%d|%s|%d", method, paging.Offset, paging.Limit, paging.SortKey, paging.SortVal)
}

func (c *CachingManager) AcademicDegreeList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "AcademicDegreeList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.AcademicDegreeList(ctx, paging)
	})
}

func (c *CachingManager) EducationList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "EducationList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.EducationList(ctx, paging)
	})
}

func (c *CachingManager) FacultyList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "FacultyList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.FacultyList(ctx, paging)
	})
}

func (c *CachingManager) InstitutionList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "InstitutionList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.InstitutionList(ctx, paging)
	})
}

func (c *CachingManager) NationalityList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "NationalityList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.NationalityList(ctx, paging)
	})
}

func (c *CachingManager) SexList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "SexList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.SexList(ctx, paging)
	})
}

func (c *CachingManager) SocialStatusList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.SocialStatus, error) {
	return cached(c, "SocialStatusList", paging, func() (int64, []*rcmodel.SocialStatus, error) {
		return c.next.SocialStatusList(ctx, paging)
	})
}

func (c *CachingManager) SpecialityList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return cached(c, "SpecialityList", paging, func() (int64, []*rcmodel.Dictionary, error) {
		return c.next.SpecialityList(ctx, paging)
	})
}

func (c *CachingManager) TypeCardList(ctx context.Context, paging *rcmodel.Paging) (int64, []*rcmodel.TypeCard, error) {
	return cached(c, "TypeCardList", paging, func() (int64, []*rcmodel.TypeCard, error) {
		return c.next.TypeCardList(ctx, paging)
	})
}

var _ IManager = (*CachingManager)(nil)
