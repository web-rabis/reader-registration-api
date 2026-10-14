package dictionary

import (
	"context"
	"errors"
	"testing"

	rcmodel "github.com/web-rabis/reader-client/model"
)

type countingManager struct {
	calls int
	items []*rcmodel.Dictionary
	err   error
}

func (m *countingManager) AcademicDegreeList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	m.calls++
	return int64(len(m.items)), m.items, m.err
}
func (m *countingManager) EducationList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return 0, nil, nil
}
func (m *countingManager) FacultyList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return 0, nil, nil
}
func (m *countingManager) InstitutionList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return 0, nil, nil
}
func (m *countingManager) NationalityList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return 0, nil, nil
}
func (m *countingManager) SexList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return 0, nil, nil
}
func (m *countingManager) SocialStatusList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.SocialStatus, error) {
	return 0, nil, nil
}
func (m *countingManager) SpecialityList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.Dictionary, error) {
	return 0, nil, nil
}
func (m *countingManager) TypeCardList(context.Context, *rcmodel.Paging) (int64, []*rcmodel.TypeCard, error) {
	return 0, nil, nil
}

func TestCachingManager_CachesRepeatedCalls(t *testing.T) {
	next := &countingManager{items: []*rcmodel.Dictionary{{Id: 1, Name: "a"}}}
	c := NewCachingManager(next)
	ctx := context.Background()

	for i := 0; i < 3; i++ {
		count, list, err := c.AcademicDegreeList(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		if count != 1 || len(list) != 1 {
			t.Errorf("i=%d count=%d list=%v", i, count, list)
		}
	}
	if next.calls != 1 {
		t.Errorf("должен вызвать нижележащий менеджер один раз, вызвано %d", next.calls)
	}
}

func TestCachingManager_DifferentPagingIsNotShared(t *testing.T) {
	next := &countingManager{items: []*rcmodel.Dictionary{{Id: 1}}}
	c := NewCachingManager(next)
	ctx := context.Background()

	if _, _, err := c.AcademicDegreeList(ctx, &rcmodel.Paging{Limit: 10}); err != nil {
		t.Fatal(err)
	}
	if _, _, err := c.AcademicDegreeList(ctx, &rcmodel.Paging{Limit: 20}); err != nil {
		t.Fatal(err)
	}
	if next.calls != 2 {
		t.Errorf("разные paging должны считаться разными ключами, вызвано %d", next.calls)
	}
}

func TestCachingManager_DoesNotCacheErrors(t *testing.T) {
	boom := errors.New("boom")
	next := &countingManager{err: boom}
	c := NewCachingManager(next)
	ctx := context.Background()

	if _, _, err := c.AcademicDegreeList(ctx, nil); !errors.Is(err, boom) {
		t.Fatalf("ожидалась ошибка: %v", err)
	}
	if _, _, err := c.AcademicDegreeList(ctx, nil); !errors.Is(err, boom) {
		t.Fatalf("ожидалась ошибка: %v", err)
	}
	if next.calls != 2 {
		t.Errorf("ошибки не должны кешироваться, вызвано %d", next.calls)
	}
}
