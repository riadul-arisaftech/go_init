package seeder

func (s *Seed) UserSeed() {
	//var roles []*model.RoleTbl
	//s.db.Find(&roles)
	//roleCount := len(roles)
	//
	//superAdmin := model.UserTbl{
	//	ID:             faker.UUIDHyphenated(),
	//	Name:           method.Ptr("Admin"),
	//	Email:          "admin@admin.com",
	//	ProfilePicture: nil,
	//	IsVerified:     method.Ptr(true),
	//	Password:       method.Ptr(str.Hash("123123")),
	//	BaseTimestamps: model.BaseTimestamps{},
	//}
	//s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&superAdmin)
	//
	//err := assignAdminRoleToAdminUser(s.db, superAdmin.ID)
	//if err != nil {
	//	return
	//}
	//
	//for i := 0; i < 50; i++ {
	//	// Generate Users
	//	userID := faker.UUIDHyphenated()
	//	err := s.db.Exec(`
	//		INSERT INTO users (id, name, email, profile_picture, is_verified) VALUES
	//		('` + userID + `',
	//			'` + faker.Name() + `',
	//			'` + faker.Email() + `',
	//			'https://i.ytimg.com/vi/gUIJ-UkQsXI/maxresdefault.jpg',
	//			'true'
	//		)`).Error
	//	if err != nil {
	//		panic(err)
	//	}
	//	// Generate role_users
	//	roleID := faker.UUIDHyphenated()
	//
	//	randIdx := rand.Intn(roleCount)
	//	err = s.db.Exec(`
	//		INSERT INTO role_users (id, user_id, role_id) VALUES
	//		('` + roleID + `',
	//			'` + userID + `',
	//			'` + roles[randIdx].ID + `'
	//		)`).Error
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}

//func assignAdminRoleToAdminUser(db *gorm.DB, id string) error {
//	var role model.Role
//	var roleUser model.RoleUser
//	if err := db.Where("name = 'Admin'").Find(&role).Error; err != nil {
//		return err
//	}
//
//	roleUser.UserID = id
//	roleUser.RoleID = role.ID
//	if err := db.Save(&roleUser.RoleUserTbl).Error; err != nil {
//		return err
//	}
//	return nil
//}
