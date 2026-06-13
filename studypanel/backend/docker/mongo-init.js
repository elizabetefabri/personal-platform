db = db.getSiblingDB("studypanel");

db.createCollection("study_items");

db.study_items.createIndex({ section: 1 });
db.study_items.createIndex({ topic: 1 });
db.study_items.createIndex({ section: 1, topic: 1 });
db.study_items.createIndex({ status: 1 });
db.study_items.createIndex({ created_at: -1 });

db.study_items.insertMany([
  {
    section: "cloud",
    topic: "aws",
    course_name: "AWS Certified Solutions Architect Associate SAA-C03",
    status: "Em andamento",
    date: "2024-06-13",
    url: "https://udemy.com/course/aws-saa",
    image_url: "",
    detail_route: "/cloud/aws",
    created_at: new Date(),
    updated_at: new Date()
  },
  {
    section: "frontend",
    topic: "angular",
    course_name: "Angular Avançado com TypeScript",
    status: "Concluído",
    date: "2024-05-01",
    url: "https://udemy.com/course/angular",
    image_url: "",
    detail_route: "/frontend/angular",
    created_at: new Date(),
    updated_at: new Date()
  },
  {
    section: "devops",
    topic: "terraform",
    course_name: "Terraform do Zero ao Avançado",
    status: "Não iniciado",
    date: "",
    url: "https://udemy.com/course/terraform",
    image_url: "",
    detail_route: "/devops/terraform",
    created_at: new Date(),
    updated_at: new Date()
  }
]);

print("MongoDB StudyPanel inicializado com sucesso.");
