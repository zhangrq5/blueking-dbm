# Generated by Django 3.2.25 on 2025-01-09 10:05

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("db_meta", "0044_deviceclass"),
    ]

    operations = [
        migrations.AlterField(
            model_name="polarisentrydetail",
            name="polaris_name",
            field=models.CharField(default="", max_length=270, unique=True),
        ),
    ]