package seeds

import "gorm.io/gorm"

func SeedTopics(db *gorm.DB) error {
	sql := `
	INSERT INTO topics (id, name, description, created_at, updated_at)
	VALUES
		(gen_random_uuid(), 'announcements', '', NOW(), NOW()),
		(gen_random_uuid(), 'Art', '', NOW(), NOW()),
		(gen_random_uuid(), 'AskReddit', '', NOW(), NOW()),
		(gen_random_uuid(), 'askscience', '', NOW(), NOW()),
		(gen_random_uuid(), 'aww', '', NOW(), NOW()),
		(gen_random_uuid(), 'blog', '', NOW(), NOW()),
		(gen_random_uuid(), 'books', '', NOW(), NOW()),
		(gen_random_uuid(), 'creepy', '', NOW(), NOW()),
		(gen_random_uuid(), 'dataisbeautiful', '', NOW(), NOW()),
		(gen_random_uuid(), 'DIY', '', NOW(), NOW()),
		(gen_random_uuid(), 'Documentaries', '', NOW(), NOW()),
		(gen_random_uuid(), 'EarthPorn', '', NOW(), NOW()),
		(gen_random_uuid(), 'explainlikeimfive', '', NOW(), NOW()),
		(gen_random_uuid(), 'food', '', NOW(), NOW()),
		(gen_random_uuid(), 'funny', '', NOW(), NOW()),
		(gen_random_uuid(), 'Futurology', '', NOW(), NOW()),
		(gen_random_uuid(), 'gadgets', '', NOW(), NOW()),
		(gen_random_uuid(), 'gaming', '', NOW(), NOW()),
		(gen_random_uuid(), 'GetMotivated', '', NOW(), NOW()),
		(gen_random_uuid(), 'gifs', '', NOW(), NOW()),
		(gen_random_uuid(), 'history', '', NOW(), NOW()),
		(gen_random_uuid(), 'IAmA', '', NOW(), NOW()),
		(gen_random_uuid(), 'InternetIsBeautiful', '', NOW(), NOW()),
		(gen_random_uuid(), 'Jokes', '', NOW(), NOW()),
		(gen_random_uuid(), 'LifeProTips', '', NOW(), NOW()),
		(gen_random_uuid(), 'listentothis', '', NOW(), NOW()),
		(gen_random_uuid(), 'mildlyinteresting', '', NOW(), NOW()),
		(gen_random_uuid(), 'movies', '', NOW(), NOW()),
		(gen_random_uuid(), 'Music', '', NOW(), NOW()),
		(gen_random_uuid(), 'news', '', NOW(), NOW()),
		(gen_random_uuid(), 'nosleep', '', NOW(), NOW()),
		(gen_random_uuid(), 'nottheonion', '', NOW(), NOW()),
		(gen_random_uuid(), 'OldSchoolCool', '', NOW(), NOW()),
		(gen_random_uuid(), 'personalfinance', '', NOW(), NOW()),
		(gen_random_uuid(), 'philosophy', '', NOW(), NOW()),
		(gen_random_uuid(), 'photoshopbattles', '', NOW(), NOW()),
		(gen_random_uuid(), 'pics', '', NOW(), NOW()),
		(gen_random_uuid(), 'science', '', NOW(), NOW()),
		(gen_random_uuid(), 'Showerthoughts', '', NOW(), NOW()),
		(gen_random_uuid(), 'space', '', NOW(), NOW()),
		(gen_random_uuid(), 'sports', '', NOW(), NOW()),
		(gen_random_uuid(), 'television', '', NOW(), NOW()),
		(gen_random_uuid(), 'tifu', '', NOW(), NOW()),
		(gen_random_uuid(), 'todayilearned', '', NOW(), NOW()),
		(gen_random_uuid(), 'UpliftingNews', '', NOW(), NOW()),
		(gen_random_uuid(), 'videos', '', NOW(), NOW()),
		(gen_random_uuid(), 'worldnews', '', NOW(), NOW())
	ON CONFLICT (name) DO NOTHING;
	`
	return db.Exec(sql).Error
}
